# HTTP

![HTTP Request & Response](http://static.ddmcdn.com/gif/dns-rev-1.gif)

The Hypertext Transfer Protocol (HTTP) is an application protocol for distributed, collaborative, hypermedia information systems.[1] HTTP is the foundation of data communication for the World Wide Web.

Hypertext is structured text that uses logical links (hyperlinks) between nodes containing text. HTTP is the protocol to exchange or transfer hypertext.

## URL

```
http://google.com
```

A uniform resource locator, abbreviated URL (also known as web address, particularly when used with HTTP), is a specific character string that constitutes a reference to a resource. In most web browsers, the URL of a web page is displayed on top inside an address bar. A URL is technically a type of uniform resource identifier (URI), but in many technical documents and verbal discussions, URL is often used as a synonym for URI, and this is not considered a problem. URLs are commonly used for web pages (http), but can also be used for file transfer (ftp), email (mailto), telephone numbers (tel) and many other applications.

## Request
```
GET /index.html HTTP/1.1
Host: www.example.com
```

The request message consists of the following:  
- A request line, for example GET /images/logo.png HTTP/1.1, which requests a resource called /images/logo.png from the server.  
- Request Headers, such as Accept-Language: en  
- An empty line.  
- An optional message body.

## Response

```
HTTP/1.1 200 OK
Date: Mon, 23 May 2005 22:38:34 GMT
Server: Apache/1.3.3.7 (Unix) (Red-Hat/Linux)
Last-Modified: Wed, 08 Jan 2003 23:11:55 GMT
ETag: "3f80f-1b6-3e1cb03b"
Content-Type: text/html; charset=UTF-8
Content-Length: 131
Connection: close

<html>
<head>
  <title>An Example Page</title>
</head>
<body>
  Hello World, this is a very simple HTML document.
</body>
</html>
```
The response message consists of the following:  
-  A Status-Line (for example HTTP/1.1 200 OK, which indicates that the client's request succeeded)  
- Response Headers, such as Content-Type: text/html  
- An empty line  
- An optional message body

## Status Codes

The following is a list of Hypertext Transfer Protocol (HTTP) response status codes. The first digit of the status code specifies one of five classes of response; the bare minimum for an HTTP client is that it recognises these five classes. 

* **1xx** Informational

* **2xx** Success

* **3xx** Redirection

* **4xx** Client Error

* **5xx** Server Error

----

## API

An application programming interface (API) specifies how some software components should interact with each other.

In addition to accessing databases or computer hardware, such as hard disk drives or video cards, an API can be used to ease the work of programming graphical user interface components. In practice, many times an API comes in the form of a library that includes specifications for routines, data structures, object classes, and variables. 

In some other cases, notably for SOAP and REST services, an API comes as just a specification of remote calls exposed to the API consumers.

### Web API

When used in the context of web development, an API is typically defined as a set of Hypertext Transfer Protocol (HTTP) request messages, along with a definition of the structure of response messages, which is usually in an Extensible Markup Language (XML) or JavaScript Object Notation (JSON) format. While "web API" historically has been virtually synonymous for web service, the recent trend (so-called Web 2.0) has been moving away from Simple Object Access Protocol (SOAP) based web services and service-oriented architecture (SOA) towards more direct representational state transfer (REST) style web resources and resource-oriented architecture (ROA). Part of this trend is related to the Semantic Web movement toward Resource Description Framework (RDF), a concept to promote web-based ontology engineering technologies. Web APIs allow the combination of multiple APIs into new applications known as mashups.

## Using APIs With Ruby

Packages with Ruby are packages that are loaded by default. Two of them, open-uri and json, are not loaded automatically when running Ruby locally. In order to load them, you can use the require keyword followed by the name of the package, like so:

```
require "open-uri"  
require "json"
```

Now with the packages loaded, you can make HTTP requests, load the HTTP response, and parse the response.

```
url = "http://www.reddit.com/.json"
request_body = open(url).read

json_object = JSON.parse(request_body)
```

In the above example, the open method that the url is passed to is part of the open-uri package. It takes a url in string format, performs a GET request to the url, and returns the response. Calling read on this method returns the body of the response.

The parse method is part of the json package. This method takes a string that contains a JSON object and converts it to a Ruby hash.

```
response = "{ 'catSound':  'meow' }"
JSON.parse(response) // { catSound: 'meow' }
```