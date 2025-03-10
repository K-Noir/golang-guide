package art_asset

import (
	"context"
	"fmt"
	"github.com/mao888/go-utils/constants"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/art_asset/bean"
	"go.mongodb.org/mongo-driver/bson"
)

func RunArtAsset2() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("activelibraries")
	collGames := db.Collection("games")

	// 2、从mongo查询数据
	mActiveLibrary := make([]*bean.MActiveLibrary, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mActiveLibrary)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mActiveLibrary)

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
	for _, source := range mActiveLibrary {

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
			Type:     constants.NumberTwo,
			AuthorID: source.Author,
			//Desc:       "",
			Name:    fmt.Sprintf("%s"+" "+"%s", source.Name, source.Desc),
			MainURL: source.Url,
			UeURL:   source.UeDownloadUrl,
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
	}
}
