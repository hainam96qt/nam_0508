package request

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

func DecodeJSON(ctx context.Context, body io.ReadCloser, dst interface{}) error {
	if err := json.NewDecoder(body).Decode(dst); err != nil {
		return fmt.Errorf("failed to decode JSON request body: err=%s", err)
	}
	defer body.Close()
	return nil
}
