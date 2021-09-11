package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"yuki0920/go-blog/model"
	"yuki0920/go-blog/repository"

	"github.com/labstack/echo/v4"
)

// ハンドラ関数という MVCにおけるコントローラーのアクションの位置づけ
func ArticleIndex(c echo.Context) error {
	if c.Request().URL.Path == "/articles" {
		c.Redirect(http.StatusPermanentRedirect, "/")
	}

	articles, err := repository.ArticleListByCursor(0)

	if err != nil {
		c.Logger().Error(err.Error())

		return c.NoContent(http.StatusInternalServerError)
	}

	// 取得できた最後の記事の ID をカーソルとして設定
	var cursor int
	if len(articles) != 0 {
		cursor = articles[len(articles)-1].ID
	}

	data := map[string]interface{}{
		"Articles": articles,
		"Cursor":   cursor,
	}

	return render(c, "article/index.html", data)
}

func ArticleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article New",
		"Now":     time.Now(),
	}

	return render(c, "article/new.html", data)
}

func ArticleShowData(c echo.Context) error {
	// パスパラメータから記事 ID を取得
	id, _ := strconv.Atoi(c.Param("articleID"))

	article, err := repository.ArticleGetByID(id)
	if err != nil {
		c.Logger().Error(err.Error())

		return c.NoContent(http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"Article": article,
	}

	return render(c, "article/show.html", data)
}

func ArticleShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("articleID"))

	article, err := repository.ArticleGetByID(id)
	if err != nil {
		c.Logger().Error(err.Error())

		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, article)
}

func ArticleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("articleID"))

	article, err := repository.ArticleGetByID(id)
	if err != nil {
		c.Logger().Error(err.Error())

		return c.NoContent(http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"Article": article,
	}

	return render(c, "article/edit.html", data)
}

type ArticleCreateOutput struct {
	Article          *model.Article
	Message          string
	ValidationErrors []string
}

func ArticleCreate(c echo.Context) error {
	// 送信されてくるフォームの内容を格納する構造体を宣言
	var article model.Article

	// レスポンスとして返却する構造体を宣言
	var out ArticleCreateOutput

	// フォームの内容を構造体にバインド
	if err := c.Bind(&article); err != nil {
		c.Logger().Error(err.Error())

		return c.JSON(http.StatusBadRequest, out)
	}

	if err := c.Validate(&article); err != nil {
		c.Logger().Error(err.Error())

		out.ValidationErrors = article.ValidationErrors(err)

		return c.JSON(http.StatusUnprocessableEntity, out)
	}

	// repository を呼び出して保存処理を実行
	res, err := repository.ArticleCreate(&article)
	if err != nil {
		c.Logger().Error(err.Error())

		return c.JSON(http.StatusInternalServerError, out)
	}

	// SQL 実行結果から作成されたレコードの ID を取得
	id, _ := res.LastInsertId()

	// 構造体に ID をセット
	article.ID = int(id)

	// レスポンスの構造体に保存した記事のデータを格納
	out.Article = &article

	// JSONにパースしてレスポンスを返却
	return c.JSON(http.StatusOK, out)
}

func ArticleDelete(c echo.Context) error {
	// パスパラメータから記事 ID を取得
	// 文字列型で取得されるので、strconv パッケージを利用して数値型にキャスト
	id, _ := strconv.Atoi(c.Param("articleID"))

	if err := repository.ArticleDelete(id); err != nil {
		c.Logger().Error(err.Error())

		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("Article %d is deleted.", id))
}

func ArticleList(c echo.Context) error {
	// 文字列型で取得できるので strconv パッケージを用いて数値型にキャスト
	cursor, _ := strconv.Atoi(c.QueryParam("cursor"))

	// 引数にカーソルの値を渡して、ID のどの位置から 10 件取得するかを指定
	articles, err := repository.ArticleListByCursor(cursor)

	if err != nil {
		c.Logger().Error(err.Error())

		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, articles)
}

type ArticleUpdateOutput struct {
	Article          *model.Article
	Message          string
	ValidationErrors []string
}

// ArticleUpdate ...
func ArticleUpdate(c echo.Context) error {
	ref := c.Request().Referer()

	refID := strings.Split(ref, "/")[4]

	reqID := c.Param("articleID")

	if reqID != refID {
		return c.JSON(http.StatusBadRequest, "")
	}

	var article model.Article
	var out ArticleUpdateOutput

	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusBadRequest, out)
	}

	if err := c.Validate(&article); err != nil {
		out.ValidationErrors = article.ValidationErrors(err)
		return c.JSON(http.StatusUnprocessableEntity, out)
	}

	articleID, _ := strconv.Atoi(reqID)
	article.ID = articleID

	_, err := repository.ArticleUpdate(&article)
	if err != nil {
		out.Message = err.Error()

		return c.JSON(http.StatusInternalServerError, out)
	}

	out.Article = &article

	return c.JSON(http.StatusOK, out)
}

func ArticleSample(c echo.Context) error {
	var article model.Article

	article.ID = 1
	article.Title = "Sample Article"
	article.Body = "Sample Article Body"

	return c.JSON(http.StatusOK, article)
}