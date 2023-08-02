package types

type Ambient struct{
    Category string
    ID string;
    Location struct{
        piso int;
        geometría string;
        pavilion string;
    };
    Description string;
    Fotos []string;
    Tags []string;
}

