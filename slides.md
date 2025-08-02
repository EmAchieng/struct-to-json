---
marp: true
theme: gaia
backgroundColor: #f9f6f2
color: #2c3e3e
---

# Structs to JSON: How Go Powers REST APIs

##### A comprehensive exploration of Go's powerful struct-to-JSON capabilities for building robust RESTful services


<span style="
  font-size:1em;
  color:#5d6d7e;
  display: block; /* Make it a block element to control its width */
  width: 100%; /* Make it span the full width of the slide's content area */
  position: relative; /* **THIS IS CRITICAL for absolute positioning of its child** */
  padding-right: 60px; /* IMPORTANT: Create space for the image to avoid text overlap */
  box-sizing: border-box; /* Ensures padding is included in the width */
">
  by Emily Achieng
  DevOps & Software Engineer
  <!-- <img src="emily.png" alt="Emily Achieng" width="50" height="50" style="
    vertical-align:middle; /* This will be overridden by absolute positioning */
    margin-left:8px; /* This will be overridden by absolute positioning */
    position: absolute; /* Take the image out of the normal flow */
    right: 0; /* Position at the right end of the *relative* parent */
    top: 50%; /* Start from the middle vertically */
    transform: translateY(-50%); /* Shift up by half its height for perfect vertical centering */
  "> -->
</span>

---

<style>
/* Custom styles for the cover slide */
section.cover-slide {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
  height: 100%;
  padding: 80px;
}

.cover-slide h1 {
  font-size: 3.5em;
  margin-bottom: 0.4em;
  line-height: 1.1;
  color: #2c3e3e;
  max-width: 90%;
}

.cover-slide h3 {
  font-size: 1.2em;
  margin-top: 0;
  max-width: 70%;
  margin-left: auto;
  margin-right: auto;
  color: #5d6d7e;
}


</style>


## Agenda

**Foundations**  
Go structs fundamentals, JSON as API language

**Implementation**  
JSON Serialization

**Advanced Techniques**  
Custom marshaling, validation

---

## Go Structs: Building Blocks of Data

Go structs are collections of fields with typed data 

- **Custom types**
- **Grouping data**
- **Object-oriented patterns**
- **Type safety**

Think of structs as **blueprints** for your data.

---

## Example: Go Struct Definition
Here’s a typical Go struct used for user data in REST APIs:
```go
type User struct {
  ID        int    `json:"id"`
  Username  string `json:"username"`
  Email     string `json:"email"`
  Active    bool   `json:"active"`
  CreatedAt time.Time `json:"created_at"`
}
```

---

## Structs to JSON: The Transformation

- JSON is the standard for data exchange in web APIs.
- Go's <span style="color:green;">encoding/json</span> package makes it simple to convert between Go structs and JSON.
- The main function for this is <span style="color:green;">json.Marshal</span>, which turns a Go struct into its JSON text representation.

---

## Example: Struct to JSON Conversion

Here’s how you convert a Go struct to JSON using `json.Marshal`:

```go
import "encoding/json"
type Product struct {
  Name string
  Price float64
}
p := Product{Name: "Laptop", Price: 1200.50}
jsonData, _ := json.Marshal(p)
// Output: {"Name":"Laptop","Price":1200.5}
```
---

## Structs to JSON: The Transformation


**1. Go Structs**  
Strongly typed data structures with fields and methods

**2. Struct Tags**  
Metadata for controlling serialization behavior

**3. Encoding/Decoding**  
`json.Marshal()` and `json.Unmarshal()` functions

**4. JSON Response**  
Client-consumable data format

---

## Struct Tags: The Secret Sauce

Struct tags are string literals that attach metadata to fields:

```go
type User struct {  
  ID int `db:"id" json:"id"`  
  Username string `db:"username" json:"username"`  
  Email string `db:"email" json:"email"`  
  Active bool `db:"active" json:"active"`  
  CreatedAt time.Time `db:"created_at" json:"created_at"`  
}
```

---

## Common JSON Tag Options

<span style="font-size:0.95em;">**1. `json:"name"`**</span>  
Rename field (change the JSON key name for this field)

<span style="font-size:0.95em;">**2. `json:"name,omitempty"`**</span>  
Skip empty values (omit the field from JSON if it is empty)

These options help you control how your struct fields are represented in JSON, making your API responses cleaner and more flexible.

---

## Common JSON Tag Options (cont.)

<span style="font-size:0.95em;">**3. `json:"-"`**</span>  
Exclude from JSON (do not include this field in the output)

<span style="font-size:0.95em;">**4. `json:",string"`**</span>  
Force string encoding (convert the field to a JSON string)

The `json:"fieldname"` tag controls how the field appears in JSON output.

---

## Request, Process, Decode, Encode

- **Request**  
  HTTP request arrives with JSON payload  
  `{ "username": "gopher", "email": "go@example.com", "active": true }`

- **Process Data**  
  Perform operations, such as validation and database storage, using the typed Go struct  
  `if err := user.Validate(); err != nil { // Handle validation error }`  
  `createdUser, err := store.CreateUser(user)`

---

## Request, Process, Decode, Encode (cont.)

- **Decode**  
  Parse JSON into struct  
  `var user model.User`  
  `json.NewDecoder(r.Body).Decode(&user)`

- **Encode**  
  Convert struct back to JSON  
  `w.Header().Set("Content-Type", "application/json")`  
  `json.NewEncoder(w).Encode(createdUser)`

---

## Validation: Ensuring Data Integrity

- **What is <span style="color:green;">Validation</span>?**  
  Ensures user data follows business rules before processing.

- **Why is <span style="color:green;">Validation Important</span>?**  
  Catches bad data early and keeps the application reliable.

- **How Does <span style="color:green;">Validation Work</span>?**  
  - Define rules (e.g., <span style="color:green;">username length</span>, <span style="color:green;">email format</span>)
  - Return clear error messages if validation fails

---

## Database Integration: Struct to Row and Back

1. Define Model  
2. Execute Query  
3. Return JSON

---

## Database Integration: Example Code

```go
type User struct { ID int; Name string }
user := User{}
db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
json, _ := json.Marshal(user)
w.Write(json)
```

Use struct tags to map fields to database columns.

---

## Putting It All Together

- **Define Data Models**  
  Create structs with appropriate JSON and DB tags

- **Implement Storage Layer**  
  Database operations that convert between structs and DB rows

---

## Putting It All Together (cont.)

- **Build HTTP Handlers**  
  Process requests and encode/decode JSON

- **Register Routes**  
  Connect HTTP endpoints to handlers

Go's struct system provides a seamless pipeline from HTTP request to database and back, with type safety at every step.

---

## Error Handling: Consistent API Responses

- **Robust APIs**  
 Provide clear, consistent error messages to clients. Go allows defining custom error structs that can be marshaled directly into JSON.  

- **Standardize Errors**    
Return predictable JSON formats for all error types.

```go
type ErrorResponse struct {  
  Code int `json:"code"`  
  Message string `json:"message"`  
}
```

- **Contextual Messages**  
Provide specific details about what went wrong.

- **HTTP Status Code**  
Align JSON responses with appropriate HTTP status codes.

---

## Key Takeaways

- **Go structs are your data foundation**  
  Model complex, reliable API data.

- **Struct tags give you control**  
  Customize JSON serialization easily.

- **Clean architecture matters**  
  Separate layers for maintainable, scalable APIs.

- **Validation & error handling ensure reliability**  
  Catch bad data and provide clear feedback.

---
## Thank You!

Questions?  
Feel free to connect:  
<a href="https://www.linkedin.com/in/emmilliarchi/">
  <img src="data:image/svg+xml;base64,PHN2ZyByb2xlPSJpbWciIHZpZXdCb3g9IjAgMCAyNCAyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48dGl0bGU+TGlua2VkSW4gY29ub3IgaWNvbjwvdGl0bGU+PHBhdGggZD0iTTIwLjQ0NyAxOC43NThWMjQuMDhoLTQuNTA3di03LjQ4MWMwLTEuNzgyLS42MzktMi45OTItMi4yNDQtMi45OTItMS4yMjcgMC0xLjk1Mi44NzUtMi4yNzkgMS43MjUtLjExNS4zMDUtMC4wNy42NC0wLjA3MS45NzV2Ny43NjNoLTQuNTA3VjkuMDI1aDQuNTA3djEuOTJjLjc0MS0xLjIxMSAxLjY5Ni0yLjE0MiAzLjcwOC0yLjE0MiAyLjY4OSAwIDQuNzA3IDEuODU1IDQuNzA3IDUuODYyek01LjAyNiAwQzIuMjExIDAgMCAyLjEyNyAwIDQuNzI3UzIuMTEgNC43MjYgNS4wMjYgNC43MjYgNC43MjYtMi4xMjUgNS4wMjYtNC43MjZTNy45NDIgMCA1LjAyNiAwWiIvPjwvc3ZnPg==" alt="LinkedIn Icon" width="30" height="30"> linkedin.com/in/emmilliarchi/
</a>
<br>
<a href="https://github.com/EmAchieng">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/github/github-original.svg" alt="GitHub Icon" width="30" height="30"> github.com/EmAchieng
</a>


---