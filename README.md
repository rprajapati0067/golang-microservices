# golang-microservices
Guide to golang-microservices



Hello, and thank you for your interest in applying to Finchat for the position of Backend Go/ Go lang Software Engineer


 we are always at the cutting edge to strive to develop the best solutions to todayâs financial problems. 
 As such, this technical assignment will be used to gauge your understanding of the following backend fundamentals:

REST APIs
GET/POST requests
Working with third party APIs
Golang/programming fundamentals
Adjusting to ambiguity
Attention to detail

We expect production grade code, so your ability to address edge cases will be assessed heavily.

Problem

 handling usersâ payments and all of the edgecases that follow their execution is commonplace.
 As such, for this assignment, you will be developing your own locally-hosted REST API 
 using the Golang Fiber library (https://github.com/gofiber/fiber) in order to facilitate basic payments using the
 Stripe Golang library (https://stripe.com/docs/api). 

Your API will be hosting 3 endpoints:

POST /api/customer
POST /api/payments
GET /api/payments/:customer_id

By the end of the assignment, your API will be able to create users, allow users to make payments,
 and retrieve a history of these payments.

Setup
Create an account on Stripeâs website (https://stripe.com/) that can be used for this assignment 
(please avoid using any personal/professional accounts if you already have them on the platform). 
In this account, you will be using the test private secret key (eg. "sk_test_your_key") 
in order to authenticate your API when making calls to the Stripe API.
Download the latest version of Golang (if you donât already have it) on your local machine. Version 1.14 or higher is required.
Initialize your Golang project and create a main.go file
Create a âHello Worldâ function (or something similar) to ensure your project works.
 If youâre able to run go run main.go and print out âHello Worldâ in your console, youâre all set.


Endpoints Functionality
Initialize the 3 endpoints above using the Golang Fiber library.
 Feel free to play around with the library until youâve gotten the gist for how Fiber works. 
 We at Finchat typically use Postman (https://www.postman.com/) to manually test our endpoints, but feel free to use another software to test your code.

Below, you will find the parameters and output for each of the endpoints.


POST /api/customer
Create a Stripe customer

Request body:
{
	email: string,
	stripeCreditCardToken: string // ex. tok_1JRJDsDvMUJnI35tHkK2U11L 
}

Response body:
{
	stripeCustomerID: string // ex. cus_wefnifjnijfn
}

When your API takes in the following request, it should create a new Stripe customer in your test dashboard. We will use email to easily identify the use, and the stripeCreditCardToken will be a token representation of the userâs default payment method. You will need to pass in both of these into Stripeâs Customer object in order to successfully create the customer. You can find the docs to Stripeâs Customer API here: https://stripe.com/docs/api/customers

In order to get the stripeCreditCardToken, you will need to manually generate a token yourself. You can do so by using Stripeâs token API: https://stripe.com/docs/api/tokens/create_card

Normally, we would have another entity generate the stripeCreditCardToken for us (eg. a mobile app, another HTTP endpoint, etc.), but for the simplicity sake of this assignment, for each call to this endpoint, you will manually generate a test token and pass it into the request.

Please only use the test card numbers that Stripe provides!



POST /api/payments
Create a payment for a given customer

Request body:
{
	stripeCustomerID: string, // ex. cus_jhdbjwfbkef
	amount: int
}

Response body:
{
	paymentIntentID: string // ex. pi_wefnifjnijfn
}

For this endpoint, we will be using the Stripe PaymentIntents API to create a payment for the user (https://stripe.com/docs/api/payment_intents). Make sure that the capture_method parameter for the payment intent object is set to âautomaticâ, but this is the default, so you shouldnât need to play around with this setting. Also, for currency, make sure you set it to âUSDâ for all payments for simplicity sake. All payment method types should be of type âcardâ.

GET /api/payments/:customer_id
Return all payments for a given customer

Request body:
{
	stripeCustomerID: string, // ex. cus_jhdbjwfbkef
}

Response body:
{
	payments: [
{
      "id": string, // ex. "pi_3JQGl3DvMUJnI35t0lbO0jiY"
      "amount": int
}
	]
}

For this endpoint, you will be returning all of the payments a user has made throughout their entire existence. Do not worry about pagination; for this assignment, it is sufficient to simply include all payments in one call to your API.

Host your new payments API on any platform you want! AWS, GCP, and Heroku all offer free credits for hosting APIs,
 so if you do go with this option, make sure to create a new account on these platforms if you donât have enough free credits. 

Please include detailed instructions for how to send HTTP requests to your API from Postman. 

Turning in the Assignment
Congratulations on completing the assignment! We appreciate your time and hope you had fun along the way. 

How to run your API
Noteable design decisions you made
Description of any outstanding 3rd party integrations you may have used
How long it took you to complete the base assignment (extra credit not included)

When you have finished writing the READ.me, 
  please include the Stripe secret key you used to make your HTTP requests to Stripe.
 We know this isnât a typically secure means of sending keys and wouldnât normally recommend using this method,
 but there shouldnât be anything valuable on the Stripe account you made.



