# Nepal-Address-API
Get addresses by province and district, quickly and accurately. Ideal for developers and users seeking specific address data.


# API Documentation

This documentation provides information about the available API endpoints and their usage.

## Get Provinces

- **Endpoint:** `/provinces`
- **HTTP Method:** GET
- **Description:** Retrieve a list of provinces.

#### Example Request:

```http
GET /provinces
```

#### Example Response:

```json
{
    "provinces": [
        "bagmati",
        "sudurpaschim",
        "lumbini",
        "pradesh-1",
        "madhesh",
        "gandaki",
        "karnali"
    ]
}
```

## Get Districts by Province

- **Endpoint:** `/districts/:province`
- **HTTP Method:** GET
- **Description:** Retrieve a list of districts within a specific province.

#### Parameters:

- `province` (URL parameter) - The ID or name of the province for which you want to retrieve districts.

#### Example Request:

```http
GET /districts/bagmati
```

#### Example Response:

```json
{
    "districts": [
        "sindhuli",
        "ramechhap",
        "dolakha",
        "bhaktapur",
        "dhading",
        "kathmandu",
        "kavrepalanchok",
        "lalitpur",
        "nuwakot",
        "rasuwa",
        "sindhupalchok",
        "chitwan",
        "makwanpur"
    ]
}
```

## Get Municipals by District

- **Endpoint:** `/municipals/:district`
- **HTTP Method:** GET
- **Description:** Retrieve a list of municipalities within a specific district.

#### Parameters:

- `district` (URL parameter) - The ID or name of the district for which you want to retrieve municipalities.

#### Example Request:

```http
GET /municipals/kathmandu
```

#### Example Response:

```json
{
    "municipals": [
        "kirtipur municipality",
        "shankharapur municipality",
        "nagarjun municipality",
        "kageshwori manahora municipality",
        "dakshinkali municipality",
        "budhanilakantha municipality",
        "tarakeshwor municipality",
        "kathmandu metropolitian city",
        "tokha municipality",
        "chandragiri municipality",
        "gokarneshwor municipality"
    ]
}
```
