package pkg

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Container struct {
	trie *Trie
}

func NewContainer() *Container {
	cnt := &Container{
		trie: NewTrie(),
	}

	words := readFile()
	for _, word := range words {
		cnt.trie.Insert(word)
	}

	return cnt
}

func (c *Container) HandleIndex(ctx *fiber.Ctx) error {
	return ctx.Render("index.html", fiber.Map{
		"words": readFile(),
	})
}

type createWordRequest struct {
	Word string `form:"word"`
}

func (c *Container) HandleCreateWord(ctx *fiber.Ctx) error {
	var req createWordRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Render("form", fiber.Map{
			"error": err.Error(),
			"words": readFile(),
		})
	}

	// refactor to not use fs
	words := readFile()
	words = append(words, req.Word)
	writeFile(words)

	c.trie.Insert(req.Word)

	return ctx.Render("form", fiber.Map{
		"words": words,
	})
}

type searchWordRequest struct {
	Word string `form:"word"`
}

func (c *Container) HandleSearchWord(ctx *fiber.Ctx) error {
	var req searchWordRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Render("search-results", fiber.Map{
			"results": []string{},
		})
	}

	words := c.trie.FindWordsWithPrefix(req.Word)

	return ctx.Render("search-results", fiber.Map{
		"results": words,
	})
}

func (c *Container) HandleDeleteWord(ctx *fiber.Ctx) error {
	idx, _ := ctx.ParamsInt("index", -1)

	// refactor to not use fs
	words := readFile()

	word := words[idx]
	words = append(words[:idx], words[idx+1:]...)
	writeFile(words)

	c.trie.Remove(word)

	return ctx.Render("list", fiber.Map{
		"words": words,
	})
}

func readFile() []string {
	raw, err := os.ReadFile("./words.json")
	if err != nil {
		log.Println("failed to read file", err)
		return []string{}
	}

	var words []string
	json.Unmarshal(raw, &words)

	return words
}

func writeFile(words []string) {
	raw, err := json.Marshal(words)
	if err != nil {
		log.Println("failed to encode", err)
	}

	err = os.WriteFile("./words.json", raw, 0644)
	if err != nil {
		log.Println("failed to write to file", err)

	}
}
