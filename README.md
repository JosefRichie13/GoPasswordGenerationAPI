# Backend for a Password Generation App Gin Gonic and Go

This repo has the code for a Password Generation App Backend. 

The below REST API endpoints are exposed

* GET /generateAPassword
  
  > This endpoint allows you to generate a strong password. It requires 4 query params to be sent.<br><br>
  > The Query Parameters are
  > ```
  > ?capLetters=<NO_OF_CAPTIAL_LETTERS_IN_THE_GENERATED_PASSWORD>
  > ?smallLetters=<NO_OF_SMALL_LETTERS_IN_THE_GENERATED_PASSWORD>
  > ?numbers=<NO_OF_NUMBERS_IN_THE_GENERATED_PASSWORD>
  > ?specialChars=<NO_OF_SPECIAL_CHARACTERS_IN_THE_GENERATED_PASSWORD>
  > ```
  > The response is a password using the params sent, e.g. if `?capLetters=5&smallLetters=5&numbers=5&specialChars=5` is sent, then the response will contain a 20 character password.  

  
* GET /generateALeetPassword
  
  > This endpoint allows you to generate a Leet password. Leet passwords are words which have been replaced by similar looking non alphabets<br><br>
  > For example, a common password is _**`starwars`**_ and this password can be cracked in 0 seconds, but leeting the same password results in _**`S7@12|_|_|@12S`**_, which will take 1 million years to crack.<br><br>
  > The Query Parameter is
  > ```
  > ?word=<PASSWORD_TO_BE_CONVERTED_TO_LEET> 
  > ```
  > The response is the leeted version of the password that was sent.
  > If you want to read more about Leet, refer https://www.gamehouse.com/blog/leet-speak-cheat-sheet/ 
