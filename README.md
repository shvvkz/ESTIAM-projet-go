# Employee API

API HTTP simple écrite en **Go (net/http)** permettant de gérer des employés en mémoire.

---

## Routes disponibles

### 🔹 GET /employees

**Description**
Récupère la liste de tous les employés.

**Exemple curl**

```bash
curl http://localhost:8080/employees
```

---

### 🔹 POST /employees

**Description**
Ajoute un nouvel employé.

**JSON attendu**

```json
{
  "name": "Alice",
  "salary": 5000
}
```

**Exemple curl**

```bash
curl -X POST http://localhost:8080/employees \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","salary":5000}'
```

---

### 🔹 PUT /employees/raise

**Description**
Augmente le salaire d’un employé existant.

**Informations**

* L’employé est recherché par son `id`
* Le salaire est modifié via un **pointeur**
* Si l’employé n’existe pas → erreur 404

**JSON attendu**

```json
{
  "id": 1,
  "percent": 10
}
```

**Exemple curl**

```bash
curl -X PUT http://localhost:8080/employees/raise \
  -H "Content-Type: application/json" \
  -d '{"id":1,"percent":10}'
```

---

## Lancer le serveur

```bash
go run .
```

Le serveur démarre sur :

```
http://localhost:8080
```