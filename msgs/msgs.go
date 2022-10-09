package msgs


type QuestionIn struct {
	Tags             []string `json:"tags"`
	IsAnswered       bool     `json:"is_answered"`
	LastActivityDate int      `json:"last_activity_date"`
	CreationDate     int      `json:"creation_date"`
	QuestionID       int      `json:"question_id"`
	Link             string   `json:"link"`
	Title            string   `json:"title"`
	Body             string   `json:"body"`
	Reasons []string `json:"reasons"`
}


type Embed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
	
type DiscordMessageOut struct {
	Content string `json:"content"`
	Tts     bool   `json:"tts"`
	Embeds  []Embed `json:"embeds"`
}