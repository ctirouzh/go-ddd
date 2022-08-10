# **Service**

We have all these entities, an *aggregate*, and a *repository* for it, but it doesn’t look like an application yet. That’s why we need the next component **Service**.

A service will tie all loosely coupled repositories into a business logic that fulfills the needs of a certain domain. In the tavern case, we might have an **Order** service, responsible for chaining together repositories to perform an order. So the service will hold access to a **CustomerRepository** and a **ProductRepository**. 

A service typically holds all the **repositories** needed to perform a certain business logic flow, such as an **Order**, **Api**, or **Billing**. What’s great is that you can even have a service inside a service.