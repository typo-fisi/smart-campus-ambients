# SmartCampus API

## How to:

### Run it

```sh
git clone https://github.com/typo-fisi/smart-campus-ambients.git
go run ./src/main/main.go
```
> It will start the API at your local host on port 4000

### Deploy it

```sh
docker compose up
```

## Where to look at?

Currently, the functional endpoints work as follows:

1. `GET /api/assignments` → `Assignment[]`:

<details>
```json
[
  {
    "code": "INE002",
    "name": "PROGRAMACIÓN Y COMPUTACIÓN",
    "credits": "2.0",
    "groups": [
      {
        "ambient_id": 301,
        "professor": {
        "Name": "SALINAS AZAÑA, GILBERTO ANÍBAL",
        "Code": "089575"
      },
      "enrolled": "43",
      "schedules": []
      },
      ...
    ]
  },
  ...
]
```
</details>

2. `GET /api/assignments/{ID}` → `Assignment`:

<details>
```json
[
  {
    "code": {ID},
    "name": "PROGRAMACIÓN Y COMPUTACIÓN",
    "credits": "2.0",
    "groups": [
      {
        "ambient_id": 301,
        "professor": {
        "Name": "SALINAS AZAÑA, GILBERTO ANÍBAL",
        "Code": "089575"
      },
      "enrolled": "43",
      "schedules": []
      },
      ...
    ]
  }
]
```
</details>


3. `GET /api/ambients` → `Ambient[]`:

<details>
```json
[
  {
    "name": "Biblioteca",
    "category": "miscelanea",
    "ambient_id": 121,
    "description": "Biblioteca de la facultad...",
    "gallery": [],
    "tags": []
  },
  ...
]
```
</details>

4. `GET /api/ambients/category/{category}` → `Ambient[]` (of the same category):

<details>
```json
[
  {
    "name": "Biblioteca",
    "category": {category},
    "ambient_id": 121,
    "description": "Biblioteca de la facultad...",
    "gallery": [],
    "tags": []
  },
  ...
]
```
</details>

5. `GET /api/ambients/ambient_id?id={id}` → `Ambient`:

<details>
In general the ambients will look like this:

```json
[
  {
    "name": "Biblioteca",
    "category": {category},
    "ambient_id": 121,
    "description": "Biblioteca de la facultad...",
    "gallery": [],
    "tags": []
  }
]
```

In case the id belongs to a "salon", then it will look like this:

```json
[
  {
    "name": "Salón 101",
    "category": "salones",
    "ambient_id": 112,
    "description": "Salón del viejo pabellón de la facultad con aforo regular que se encuentra en el primer nivel.",
    "gallery": [],
    "tags": [],
    "Assignments": [
      {
        "ambient_id": 112,
        "professor": {
          "Name": "PERALTA ORTIZ, VILMA",
          "Code": "10665694"
        },
        "enrolled": "36",
        "schedules": [
          {
            "day": 2,
            "type": "T",
            "from": "1400",
            "to": "1600"
          },
          ...
        ]
      },
      ...
    ]
  }
]

```
</details>

#### Detailed types info can be found at `./src/types`

