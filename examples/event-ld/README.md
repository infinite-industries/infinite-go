# event-ld
Return the JSON-LD representation of an Infinite Industries event.

```console
$ ./event-ld 4da9761e-1dfd-488c-8f53-2a95cac188c1
{
  "name": "Intuitive Collage Class",
  "startDate": "2023-07-27T22:00:00Z",
  "endDate": "2023-07-28T00:00:00Z",
  "location": {
    "name": "Lex Center for Creative ReUse",
    "address": {
      "streetAddress": "110 Luigart Court",
      "addressLocality": "Lexington",
      "postalCode": "40508",
      "addressRegion": "Kentucky",
      "addressCountry": "US",
      "@type": "PostalAddress"
    },
    "hasMap": "https://goo.gl/maps/URvk6iu5ycyMy6Dp9",
    "identifier": "745563a4-c22d-405c-ae70-00f39eb826a9",
    "@type": "Place"
  },
  "identifier": "4da9761e-1dfd-488c-8f53-2a95cac188c1",
  "image": [
    "https://infinite-industries-event-images.s3.us-west-2.amazonaws.com/uploads/3c28f34a-7f6e-4146-a289-23f851d43544.jpg"
  ],
  "brief_description": "Learn to make collage art with Oliver!",
  "description": "\u003cp\u003e\u003cspan style=\"color: rgb(51, 51, 51);\"\u003eJoin Oliver to create artwork with the fantastically fun and relaxing technique of intuitive collage!\u003c/span\u003e\u003c/p\u003e\u003cp\u003e\u003cspan style=\"color: rgb(51, 51, 51);\"\u003eAll materials are included in the cost of the class.\u003c/span\u003e\u003c/p\u003e",
  "offers": [
    {
      "url": "https://www.simpletix.com/e/intuitive-collage-with-oliver-tickets-137879",
      "validFrom": "0001-01-01T00:00:00Z",
      "@type": "Offer"
    }
  ],
  "url": "https://infinite.industries/events/4da9761e-1dfd-488c-8f53-2a95cac188c1",
  "subjectOf": [
    {
      "url": "https://www.simpletix.com/e/intuitive-collage-with-oliver-tickets-137879",
      "name": "Ticket Link",
      "@type": "WebPage"
    }
  ],
  "updatedAt": "2023-06-28T18:39:47.925Z",
  "@type": "Event",
  "@context": "https://schema.org"
}
```
