# Punto De Venta API

Punto De Venta API es una **Interfaz de Programación de Aplicaciones** genérica diseñada para servir como plantilla para el desarrollo de una aplicación para la gestión de una tienda, comercio, bazar, quiosco o para cualquier persona que necesite de una gestión adecuada de inventarios, ventas, facturación, etc.

## Estructura

```
📁 cmd
|-- 📁 api
    |-- 📄 main.go // Punto de entrada de la aplicación.
📁 internal
|-- 📁 database
    |-- 📄 database.go // Conexión con base de datos y configuración de GORM.
|-- 📁 handlers // Controladores HTTP.
|-- 📁 models // Modelos de datos.
|-- 📁 repositories // Repositorios de datos.
    |-- 📄 base.go // Repositorio base con operaciones CRUD.
```

---

## Endpoints principales

### Productos

- GET /products — listar productos.
- GET /products/{id} — ver detalle de un producto.
- GET /products/category/{category_id} - listar productos por categoría.
- POST /products — crear producto.
- PUT /products/{id} — actualizar producto.
- DELETE /products/{id} — eliminar producto.

### Categorías

- GET /categories — listar productos.
- POST /categories — crear producto.
- PUT /categories/{id} — actualizar producto.
- DELETE /categories/{id} — eliminar producto.