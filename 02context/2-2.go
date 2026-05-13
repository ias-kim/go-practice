package context

import "net/http"

func Handle(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ID int
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != null {
		// 오류 처리
	}
	b, err := GetBook(r.Context, body.ID)
	// 남은 처리
}

// 로직 냉 context를 사용할 예정이 없는 ㅁ서드
func GetBook(ctx context.Context, id int) (*Book, error) {
	// 호출한 함수나 메서드에 context가 필요한 경우도 있다.
	rows, err := db.QueryContext(ctx, "SELECT id, name, isdn, price FROM books WHERE id=?", id)
	// 남은 처리
}