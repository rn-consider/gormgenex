// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"github.com/rn-consider/gormgenex/dao/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.ArticleTag{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.ArticleTag{}) fail: %s", err)
	}
}

func Test_articleTagQuery(t *testing.T) {
	articleTag := newArticleTag(db)
	articleTag = *articleTag.As(articleTag.TableName())
	_do := articleTag.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(articleTag.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <articles_tags> fail:", err)
		return
	}

	_, ok := articleTag.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from articleTag success")
	}

	err = _do.Create(&model.ArticleTag{})
	if err != nil {
		t.Error("create item in table <articles_tags> fail:", err)
	}

	err = _do.Save(&model.ArticleTag{})
	if err != nil {
		t.Error("create item in table <articles_tags> fail:", err)
	}

	err = _do.CreateInBatches([]*model.ArticleTag{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <articles_tags> fail:", err)
	}

	_, err = _do.Select(articleTag.ALL).Take()
	if err != nil {
		t.Error("Take() on table <articles_tags> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <articles_tags> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <articles_tags> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <articles_tags> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.ArticleTag{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <articles_tags> fail:", err)
	}

	_, err = _do.Select(articleTag.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <articles_tags> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <articles_tags> fail:", err)
	}

	_, err = _do.Select(articleTag.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <articles_tags> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <articles_tags> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <articles_tags> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <articles_tags> fail:", err)
	}

	_, err = _do.ScanByPage(&model.ArticleTag{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <articles_tags> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <articles_tags> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <articles_tags> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <articles_tags> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <articles_tags> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <articles_tags> fail:", err)
	}
}
