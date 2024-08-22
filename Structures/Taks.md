### Structures

1. **Simple structure:** Create a Person structure that contains the fields Name (string), Age (int), and Address (string). Then create a function that takes an object of this structure and prints all of its fields.

2. **Nested structure:** Create an Address structure that contains the fields Street (string), City (string), and PostalCode (string). Then nest this structure in a Person structure. Create a function that prints the full address of the person.

3. **Methods on the structure:** Add a Greet() method to the Person structure that prints a greeting message in the format: "Hello, my name is \[Name\] and I am \[Age\] years old."

4. **Changing the value of the field:** Add a HaveBirthday() method to the Person structure that increments the age of the person by 1 year. Then test this method on the object.

5. **Copying structures:** Create two Person structures, copy one to the other, change the data in the copy, and check if the data in the original structure has changed as well.

6. **Map of structures:** Create a map where the key is a string (e.g. the person's name), and the value is a Person structure. Then populate the map with several people and print their data.

7. **Anonymous fields:** Create an Employee structure that will anonymously inherit fields from the Person structure. Add an additional Position field (string) and a Describe() method that prints all the information about the employee.

8. **Initializing structures:** Write a NewPerson function that will create and return a new Person structure object. Add support for different initialization methods: with default values ​​and with values ​​provided as arguments.

9. **Comparing structures:** Create two Person structures, compare them, and print if they are identical.

10. **JSON and Structures:** Create a Person structure, serialize it to JSON, then deserialize the JSON back to a structure. Verify that the data has been transformed correctly.