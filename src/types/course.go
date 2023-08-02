package types;

/*
interface AssignmentGroup {
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
    enrolled: number,
    ambient_id: string // número de aula
}
*/

type AssignmentGroup struct {
    Ambient_id string;
    Group string;
    Professor Professor;
    Enrolled int;

    Schedule struct{
        Weekday int;

        Theory struct{
            From int;
            To int;
        };

        Practice struct{
            From int;
            To int;
        }
    };

}

