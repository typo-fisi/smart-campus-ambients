package types

type Ambient struct{
    Category string
    Id string;
    Location struct{
        piso int;
        geometría string;
        pavilion string;
    };
    Description string;
    Fotos []string;
    Tags []string;
}

