package art_asset

import (
	"context"
	"fmt"
	"github.com/mao888/go-utils/constants"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/art_asset/bean"
	"go.mongodb.org/mongo-driver/bson"
)

func RunArtAsset1() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("artsources")
	collGames := db.Collection("games")
	collCloudurls := db.Collection("cloudurls")

	// 2、从mongo查询数据
	mArtSource := make([]*bean.MArtSource, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mArtSource)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mArtSource)

	var mGame []bean.MGame
	err = collGames.Find(context.TODO(), bson.M{}).All(&mGame)
	if err != nil {
		fmt.Println("Mongo查询MGame错误：", err)
	}

	// 3、将从mongo中查出的games.id(int)作为key, games.game_id(string)作为value,存入map
	idMap := map[int32]string{}
	for _, game := range mGame {
		idMap[game.ID] = game.GameID
	}

	// 3、将mongo数据装入切片
	//artAsset := make([]*bean.ArtAsset, 0)
	for _, source := range mArtSource {

		// DoneAt
		var doneAt int64
		if source.DoneTime != nil {
			doneAt = source.DoneTime.Unix()
		}
		// UpdateTime
		var updateTime int64
		if source.UpdateTime != nil {
			updateTime = source.UpdateTime.Unix()
		}
		// CreatedAt
		var createdAt int64
		if source.CreateTime != nil {
			createdAt = source.CreateTime.Unix()
		}
		artAsset := &bean.ArtAsset{
			ID:       source.ID,
			Type:     constants.NumberOne,
			AuthorID: source.Author,
			//Desc:       "",
			Name: fmt.Sprintf("%s"+" "+"%s", source.Name, source.Desc),
			//MainURL: "",
			UeURL:   source.DownloadUrl,
			MayaURL: source.MayaDownloadUrl,
			//Remark:     "",
			GameID:     idMap[source.GameId],
			IsInternal: true,
			//CategorieID: 0,
			CreatedAt: createdAt,
			UpdatedAt: updateTime,
			DoneAt:    doneAt,
			IsDeleted: false,
		}
		// 4、将装有mongo数据的切片入库
		err = db2.MySQLClientCruiser.Table("art_asset").Create(artAsset).Error
		if err != nil {
			fmt.Println("入mysql/art_asset错误：", err)
		}

		// 将ThumbArr 预览图 插入 mysql/asset_imgs 资源库-主图表
		if len(source.ThumbArr) == constants.NumberZero {
			continue
		}
		for _, i := range source.ThumbArr {
			// 根据Id去mongo/cloudurls 查询，并插入mysql/asset_imgs
			mCloudUrls := make([]*bean.MCloudUrls, 0)
			err := collCloudurls.Find(context.TODO(), bson.M{"_id": i}).All(&mCloudUrls)
			if err != nil {
				fmt.Println("Mongo查询错误：", err)
				return
			}
			if len(mCloudUrls) == constants.NumberZero {
				continue
			}
			fmt.Printf("thumb id:%d, cloudurls:%v \n", i, mCloudUrls)
			assetImg := &bean.AssetImg{
				//ID:        i,
				AssetID:   artAsset.ID,
				Name:      mCloudUrls[0].Name,
				URL:       mCloudUrls[0].Url,
				SizeRatio: "",
				Size:      mCloudUrls[0].Size,
				Md5:       "",
				Height:    0,
				Width:     0,
				IsDeleted: false,
			}
			// 入mysql/asset_imgs
			err = db2.MySQLClientCruiser.Table("asset_imgs").Create(assetImg).Error
			if err != nil {
				fmt.Println("入mysql/asset_imgs错误：", err)
			}
		}
	}
}
