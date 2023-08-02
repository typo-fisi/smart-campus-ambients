# Schema

So the idea of this is to get the information of the classes through endpoints in this api, like so:

- `GET /teachers` -> returns:

```json
{
    [
        name: String,
        code: String,
        courses: {
            [
                {
                    name: String,
                    ID: String
                },
                ...
            ]
        }
        ...
    ]
}
```

```ts
interface Course {
    name: string,
    credits: number,
    class_code: string, // INE002
    professor: Professor,
    schedule: {
        weekday: number,
        theory: {
            from: number,
            to: number
        },
        practice: {
            from: number,
            to: number
        }
    },
    group: string, // sección
    enrolled: number
}

```

Alternatively, if we already have a certain teacher's ID -> `GET /teacher/{Code}`, which should return a single teacher's information.

We should also be able to get only the certain course's information

- `GET /course/{ambient_id}`



