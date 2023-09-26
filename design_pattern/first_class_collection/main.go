package first_class_collection

type Article struct {
	ID          string
	Title       string
	IsPublished bool
}

// ArticleList ファーストクラスコレクションとはリストやマップなどのコレクションをラップした独自の型のこと
type ArticleList []*Article

// Published 公開済みの記事を取得する
// Entityの構造体に振る舞いを持たせることで、Entityの振る舞いを変更する際に、
// そのEntityを扱うロジックを変更するだけで済むようになる
func (al ArticleList) Published() ArticleList {
	var published ArticleList
	for _, article := range al {
		if article.IsPublished {
			published = append(published, article)
		}
	}
	return published
}

func FindArticlesByUserID(userID string) ([]*Article, error) {

	return ArticleList{
		&Article{
			ID:          "1",
			Title:       "title1",
			IsPublished: true,
		},
		&Article{
			ID:          "1",
			Title:       "title2",
			IsPublished: false,
		},
	}, nil
}

// FindPublishedArticlesByUserID ユーザーIDを指定して公開済みの記事を取得する
func FindPublishedArticlesByUserID(userID string) ([]*Article, error) {

	articles, err := FindArticlesByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Usecase層でEntityの構造体を扱うロジックが散らばってしまう
	// データとその振る舞いをセットで実装できていないので、データの振る舞いを変更する際に、
	// そのデータを扱うロジックを全て変更する必要がある
	var publicArticles []*Article
	for _, article := range articles {
		if article.IsPublished {
			publicArticles = append(publicArticles, article)
		}
	}
	return publicArticles, nil
}

func FindArticlesByUserID2(userID string) (ArticleList, error) {
	return ArticleList{
		&Article{
			ID:          "1",
			Title:       "title1",
			IsPublished: true,
		},
		&Article{
			ID:          "1",
			Title:       "title2",
			IsPublished: false,
		},
	}, nil
}

func FindPublishedArticlesByUserID2(userID string) (ArticleList, error) {

	articles, err := FindArticlesByUserID2(userID)
	if err != nil {
		return nil, err
	}
	// データとその振る舞いがセットで実装されているので、実装が単純になる
	// articles構造体がPublished()という振る舞いを持っているので
	// その振る舞いを変更する際に、そのarticles構造体を扱うロジックを変更するだけで済む
	return articles.Published(), nil
}
