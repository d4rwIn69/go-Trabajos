package main

import (
	"errors"
	"fmt"
)

type Producto struct {
	id       int
	nombre   string
	precio   float64
	cantidad int
}

func (p Producto) GetID() int {
	return p.id
}

func (p Producto) GetNombre() string {
	return p.nombre
}

func (p Producto) GetPrecio() float64 {
	return p.precio
}

func (p Producto) GetCantidad() int {
	return p.cantidad
}

func (p *Producto) SetPrecio(precio float64) error {
	if precio < 0 {
		return errors.New("ERROR: el precio no puede ser negativo")
	}
	p.precio = precio
	return nil
}

func (p *Producto) SetCantidad(cantidad int) error {
	if cantidad < 0 {
		return errors.New("ERROR: la cantidad no puede ser negativa")
	}
	p.cantidad = cantidad
	return nil
}

func NewProducto(id int, nombre string, precio float64, cantidad int) (*Producto, error) {
	if id <= 0 {
		return nil, errors.New("ERROR: el ID debe ser mayor a 0")
	}
	p := &Producto{}
	p.id = id
	p.nombre = nombre

	if err := p.SetPrecio(precio); err != nil {
		return nil, err
	}
	if err := p.SetCantidad(cantidad); err != nil {
		return nil, err
	}

	return p, nil
}

func (p Producto) MostrarInfo() {
	fmt.Printf("ID: %d | Nombre: %s | Precio: $%.2f | Cantidad: %d unidades\n",
		p.id, p.nombre, p.precio, p.cantidad)
}

type Inventario struct {
	productos map[int]Producto
}

func NuevoInventario() Inventario {
	return Inventario{
		productos: make(map[int]Producto),
	}
}

func (inv *Inventario) AgregarProducto(p Producto) error {
	if _, existe := inv.productos[p.GetID()]; existe {
		return errors.New("ERROR: ya existe un producto con ese ID")
	}
	inv.productos[p.GetID()] = p
	fmt.Printf("Producto '%s' agregado correctamente.\n", p.GetNombre())
	return nil
}

func (inv *Inventario) EliminarProducto(id int) error {
	if _, existe := inv.productos[id]; !existe {
		return errors.New("ERROR: no existe un producto con ese ID")
	}
	delete(inv.productos, id)
	fmt.Printf("Producto con ID %d eliminado.\n", id)
	return nil
}

func (inv Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := inv.productos[id]
	return p, existe
}

func (inv Inventario) ListarProductos() {
	if len(inv.productos) == 0 {
		fmt.Println("El inventario está vacío.")
		return
	}
	fmt.Println("===== LISTA DE INVENTARIO =====")
	for _, p := range inv.productos {
		p.MostrarInfo()
	}
	fmt.Println("===============================")
}

func main() {
	inventario := NuevoInventario()

	p1, err := NewProducto(1, "Laptop", 850.50, 10)
	if err == nil {
		inventario.AgregarProducto(*p1)
	}

	p2, err := NewProducto(2, "Mouse", 15.99, 50)
	if err == nil {
		inventario.AgregarProducto(*p2)
	}

	fmt.Println()
	inventario.ListarProductos()

	fmt.Println("\nBuscando y actualizando el producto con ID 2...")
	if prod, encontrado := inventario.BuscarProducto(2); encontrado {
		prod.SetPrecio(12.50)
		prod.SetCantidad(45)
		inventario.productos[prod.GetID()] = prod
		fmt.Println("¡Actualizado con éxito!")
		prod.MostrarInfo()
	}

	fmt.Println()
	inventario.EliminarProducto(1)

	fmt.Println()
	inventario.ListarProductos()
}	