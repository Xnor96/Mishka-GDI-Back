package domain

type CategoriaRepository interface {
	GetAll() ([]Categoria, error)
	GetByID(id int) (*Categoria, error)
	Create(categoria *Categoria) error
	Update(categoria *Categoria) error
	Delete(id int) error
}
