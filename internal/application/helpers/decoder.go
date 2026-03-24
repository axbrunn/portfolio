package helpers

import (
	"errors"
	"net/http"

	"github.com/go-playground/form/v4"
)

var formDecoder = form.NewDecoder()

func DecodePostForm(r *http.Request, dst any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	err = formDecoder.Decode(dst, r.PostForm)
	var invalidDecoderError *form.InvalidDecoderError
	if errors.As(err, &invalidDecoderError) {
		panic(err)
	}

	return err
}
