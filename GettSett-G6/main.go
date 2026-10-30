package main

import (
	"fmt"
)

type Producto struct {
	id       int
	nombre   string
	precio   float64
	cantidad int
}

type Inventario struct {
	productos map[int]Producto
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) *Producto {
	return &Producto{
		id:       id,
		nombre:   nombre,
		precio:   precio,
		cantidad: cantidad,
	}
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

func (p *Producto) SetPrecio(nuevoPrecio float64) {
	p.precio = nuevoPrecio
}

func (p *Producto) SetCantidad(nuevaCantidad int) {
	p.cantidad = nuevaCantidad
}

func (p Producto) MostrarInfo() {
	fmt.Printf("ID: %d | Nombre: %s | Precio: $%.2f | Cantidad: %d\n",
		p.GetID(), p.GetNombre(), p.GetPrecio(), p.GetCantidad())
}

func NuevoInventario() *Inventario {
	return &Inventario{
		productos: make(map[int]Producto),
	}
}

func (inv *Inventario) AgregarProducto(p Producto) {
	if _, existe := inv.productos[p.GetID()]; existe {
		fmt.Printf("Error: ya existe un producto con el ID %d\n", p.GetID())
		return
	}
	inv.productos[p.GetID()] = p
	fmt.Printf("Producto '%s' agregado correctamente.\n", p.GetNombre())
}

func (inv *Inventario) EliminarProducto(id int) {
	if _, existe := inv.productos[id]; !existe {
		fmt.Printf("Error: no existe un producto con el ID %d\n", id)
		return
	}
	delete(inv.productos, id)
	fmt.Printf("Producto con ID %d eliminado.\n", id)
}

func (inv *Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := inv.productos[id]
	return p, existe
}

func (inv *Inventario) ListarProductos() {
	if len(inv.productos) == 0 {
		fmt.Println("El inventario está vacío.")
		return
	}

	fmt.Println("===== INVENTARIO =====")
	for _, p := range inv.productos {
		p.MostrarInfo()
	}
	fmt.Println("======================")
}

func main() {
	inv := NuevoInventario()

	p1 := NuevoProducto(1, "Laptop", 850.00, 5)
	p2 := NuevoProducto(2, "Mouse", 15.50, 30)
	p3 := NuevoProducto(3, "Teclado", 25.00, 20)

	inv.AgregarProducto(*p1)
	inv.AgregarProducto(*p2)
	inv.AgregarProducto(*p3)

	fmt.Println()
	inv.ListarProductos()

	fmt.Println()
	if p, ok := inv.BuscarProducto(2); ok {
		fmt.Print("Búsqueda exitosa -> ")
		p.MostrarInfo()
	}

	fmt.Println()
	if p, ok := inv.BuscarProducto(1); ok {
		p.SetPrecio(799.99)
		p.SetCantidad(8)
		inv.EliminarProducto(p.GetID())
		inv.AgregarProducto(p)
	}

	fmt.Println()
	inv.EliminarProducto(2)

	fmt.Println()
	inv.ListarProductos()
}