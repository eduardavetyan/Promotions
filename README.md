# Promotions

## Requirement

We receive some records in a CSV file (example promotions.csv attached) every 30
minutes. We would like to store these objects in a way to be accessed by an endpoint.
Given an ID the endpoint should return the object, otherwise, return not found.
Eg:
```bash
curl https://localhost:1321/promotions/1
{"id":"172FFC14-D229-4C93-B06B-F48B8C095512", "price":9.68,
"expiration_date": "2022-06-04 06:01:20"}
```

#### Additionally, consider:
- The .csv file could be very big (billions of entries) - how would your application
perform?
- Every new file is immutable, that is, you should erase and write the whole storage;
- How would your application perform in peak periods (millions of requests per
minute)?
- How would you operate this app in production (e.g. deployment, scaling, monitoring)?
- The application should be written in golang;
- Main deliverable is the code for the app including usage instructions, ideally in a
repo/github gist.

## Setup

1. Create new directory

```bash
mkdir promotions
cd promotions
```

2. Clone GitHub repo

```bash
git clone https://github.com/eduardavetyan/Promotions.git .
```

3. Run Docker Compose
```bash
docker compose up -d
```

## Usage

Application is now available through address: `localhost:3030`

There are 2 implemented endpoints:

1. Upload .csv file 
```bash
curl --location 'http://localhost:3030/promotions/' \
--form 'file=@"/path/to/promotions.csv"'
```

2. Get promotion item by ID
```bash
curl --location 'http://localhost:3030/promotions/d018ef0b-dbd9-48f1-ac1a-eb4d90e57118'
```