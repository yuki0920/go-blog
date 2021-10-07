package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"yuki0920/go-notes/domain/model"
	"yuki0920/go-notes/usecase"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	articleUsecase usecase.ArticleUsecase
}

func NewArticleHandler(articleUsecase usecase.ArticleUsecase) ArticleHandler {
	return ArticleHandler{
		articleUsecase: articleUsecase,
	}
}

func (handler *ArticleHandler) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("articleID"))
		article, err := handler.articleUsecase.GetById(id)
		if err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusNoContent, err)
		}

		return c.JSON(http.StatusOK, article)
	}
}

func (handler *ArticleHandler) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		var cursor int
		// 文字列型で取得できるので strconv パッケージを用いて数値型にキャスト
		cursor, _ = strconv.Atoi(c.QueryParam("cursor"))

		// 引数にカーソルの値を渡して、ID のどの位置から 10 件取得するかを指定
		articles, err := handler.articleUsecase.ListByCursor(cursor)
		if err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusInternalServerError, err)
		}

		// 取得できた最後の記事の ID をカーソルとして設定
		if len(articles) != 0 {
			cursor = articles[len(articles)-1].ID
		}

		// キーはstring,値が配列とintなのでinterface{}にしている
		data := map[string]interface{}{
			"articles": articles,
			"cursor":   cursor,
		}

		return c.JSON(http.StatusOK, data)
	}
}

type ArticleCreateOutput struct {
	Article          *model.Article
	Message          string
	ValidationErrors []string
}

func (handler *ArticleHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		// 送信されてくるフォームの内容を格納する構造体を宣言
		var article model.Article

		// レスポンスとして返却する構造体を宣言
		var out ArticleCreateOutput

		// フォームの内容を構造体にバインド
		if err := c.Bind(&article); err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusBadRequest, out)
		}

		// バインド後にバリデーションを実行
		if err := c.Validate(&article); err != nil {
			c.Logger().Error(err.Error())

			out.ValidationErrors = article.ValidationErrors(err)

			return c.JSON(http.StatusUnprocessableEntity, out)
		}

		// repository を呼び出して保存処理を実行
		id, err := handler.articleUsecase.Create(&article)
		if err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusInternalServerError, out)
		}

		// 構造体に ID をセット
		article.ID = int(id)

		// レスポンスの構造体に保存した記事のデータを格納
		out.Article = &article

		// JSONにパースしてレスポンスを返却
		return c.JSON(http.StatusOK, out)

	}
}

type ArticleUpdateOutput struct {
	Article          *model.Article
	Message          string
	ValidationErrors []string
}

func (handler *ArticleHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var out ArticleUpdateOutput
		var article model.Article

		// フォームの内容を構造体にバインドする、構造体で設定した型と異なる場合はエラーになる
		if err := c.Bind(&article); err != nil {
			out.Message = err.Error()

			return c.JSON(http.StatusBadRequest, out)
		}

		// フォームの内容を検証する
		if err := c.Validate(&article); err != nil {
			out.ValidationErrors = article.ValidationErrors(err)
			return c.JSON(http.StatusUnprocessableEntity, out)
		}

		articleID, _ := strconv.Atoi(c.Param("articleID"))
		article.ID = articleID

		err := handler.articleUsecase.Update(&article)
		if err != nil {
			out.Message = err.Error()

			return c.JSON(http.StatusInternalServerError, out)
		}

		out.Article = &article

		return c.JSON(http.StatusOK, out)
	}

}

func (handler *ArticleHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		// パスパラメータから記事 ID を取得
		// 文字列型で取得されるので、strconv パッケージを利用して数値型にキャスト
		id, _ := strconv.Atoi(c.Param("articleID"))

		if err := handler.articleUsecase.Delete(id); err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusInternalServerError, "")
		}

		return c.JSON(http.StatusOK, fmt.Sprintf("Article %d is deleted.", id))
	}
}
