package tool

import (
	"github.com/qipaiyouxi/fkmahjongserver/cache"
	"github.com/qipaiyouxi/fkmahjongserver/db"
	"github.com/qipaiyouxi/fkmahjongserver/db/dao"
	"github.com/qipaiyouxi/fkmahjongserver/def"
	"github.com/qipaiyouxi/fkmahjongserver/notice"

	log "github.com/Sirupsen/logrus"
)

func handleShopsManage(shops []*dao.Shop) error {
	for _, shop := range shops {
		daoShop := cache.GetGMTShop(shop.IndexID)
		if daoShop == nil {
			createShop(shop)
			continue
		}

		shop.SetExist(true)
		updateShop(shop)
	}

	notice.ToolInitShop()
	return nil
}

func createShop(shop *dao.Shop) {
	err := shop.Insert(db.Pool)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error(def.ErrInsertShop)
	}
}

func updateShop(shop *dao.Shop) {
	err := shop.Update(db.Pool)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error(def.ErrUpdateShop)
	}
}

func deleteShop(shop *dao.Shop) {
	err := shop.Delete(db.Pool)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error(def.ErrDeleteShop)
	}
}
