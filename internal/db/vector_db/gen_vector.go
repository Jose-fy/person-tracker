package db


type SearchResult struct{}

type VectorDatabase interface {
    Insert(id string, vector []float32) error
    Search(vector []float32, topK int) ([]SearchResult, error)
    // ... other necessary methods ...
}


