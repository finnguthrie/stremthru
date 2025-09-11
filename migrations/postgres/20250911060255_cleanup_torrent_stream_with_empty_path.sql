-- +goose Up
-- +goose StatementBegin
DELETE FROM "public"."torrent_stream" WHERE "p" = '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
