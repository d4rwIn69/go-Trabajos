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

func NuevoProducto(id int, nombre string, precio float64, cantidad int) (Producto, error) {
	if id <= 0 {
		return Producto{}, errors.New("ERROR: el ID debe ser mayor a 0")
	}
	if precio < 0 {
		return Producto{}, errors.New("ERROR: el precio no puede ser negativo")
	}
	if cantidad < 0 {
		return Producto{}, errors.New("ERROR: la cantidad no puede ser negativa")
	}

	return Producto{
		id:       id,
		nombre:   nombre,
		precio:   precio,
		cantidad: cantidad,
	}, nil
}

// Getters y Setters para Producto
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

func (p *Producto) SetPrecio(nuevoPrecio float64) error {
	if nuevoPrecio < 0 {
		return errors.New("ERROR: el precio no puede ser negativo")
	}
	p.precio = nuevoPrecio
	return nil
}

func (p *Producto) SetCantidad(nuevaCantidad int) error {
	if nuevaCantidad < 0 {
		return errors.New("ERROR: la cantidad no puede ser negativa")
	}
	p.cantidad = nuevaCantidad
	return nil
}

// MostrarInfo muestra la información completa de forma amigable
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
		return fmt.Errorf("ERROR: ya existe un producto con el ID %d", p.GetID())
	}
	inv.productos[p.GetID()] = p
	fmt.Printf("Producto '%s' agregado correctamente.\n", p.GetNombre())
	return nil
}

// EliminarProducto elimina un producto mediante su ID
func (inv *Inventario) EliminarProducto(id int) error {
	if _, existe := inv.productos[id]; !existe {
		return fmt.Errorf("ERROR: no existe un producto con el ID %d", id)
	}
	delete(inv.productos, id)
	fmt.Printf("Producto con ID %d eliminado del inventario.\n", id)
	return nil
}

// BuscarProducto busca y devuelve un producto y un booleano indicando si existe
func (inv Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := inv.productos[id]
	return p, existe
}

// ListarProductos muestra todos los productos almacenados
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

// ==========================================
// 3. FUNCIÓN MAIN (Pruebas del sistema)
// ==========================================
func main() {
	// Inicializamos el inventario
	miInventario := NuevoInventario()

	// Creamos productos usando el constructor con validación
	p1, err1 := NuevoProducto(1, "Laptop Gamer", 1200.50, 5)
	if err1 == nil {
		miInventario.AgregarProducto(p1)
	}

	p2, err2 := NuevoProducto(2, "Mouse Inalámbrico", 25.00, 30)
	if err2 == nil {
		miInventario.AgregarProducto(p2)
	}

	p3, err3 := NuevoProducto(3, "Teclado Mecánico", 75.50, 15)
	if err3 == nil {
		miInventario.AgregarProducto(p3)
	}

	fmt.Println()
	// Listar todos los productos
	miInventario.ListarProductos()

	fmt.Println("\n--- Probando Búsqueda y Modificación por Setters ---")
	// Buscar producto con ID 2
	if prodEncontrado, encontrado := miInventario.BuscarProducto(2); encontrado {
		fmt.Printf("Encontrado: %s. Precio actual: $%.2f\n", prodEncontrado.GetNombre(), prodEncontrado.GetPrecio())
		
		// Modificamos el precio usando el Setter (necesitamos un puntero o reasignar en el mapa)
		// Como map devuelve una copia, actualizamos directamente sobre el mapa usando el ID:
		// O usando un puntero si se extrajera de otra forma. Actualicemos directamente el mapa:
		prodEncontrado.SetPrecio(22.99)
		miInventario.productos[prodEncontrado.GetID()] = prodEncontrado
		
		fmt.Println("Precio actualizado exitosamente.")
	}

	fmt.Println()
	// Listar nuevamente para ver el cambio reflejado
	miInventario.ListarProductos()

	fmt.Println("\n--- Probando Eliminación ---")
	miInventario.EliminarProducto(3)

	fmt.Println()
	// Inventario final
	miInventario.ListarProductos()
}