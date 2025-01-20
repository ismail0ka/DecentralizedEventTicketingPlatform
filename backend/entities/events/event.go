package events

type Event struct {
	id                int64   `json: "id"`
	name              string  `json: "name"`
	description       string  `json: "description"`
	date              string  `json: "date"`
	location          string  `json: "location"`
	organizer_id      int64   `json: "organizer_id"`
	available_tickets int32   `json: "available_tickets"`
	price             float32 `json: "price"`
	image_url         string  `json: "image_url"`
}
