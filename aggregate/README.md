# **Aggregates**

An Aggregate is a set of entities and value objects combined. 

Martin Flower
>
> DDD aggregates are domain concepts (order, clinic visit, playlist) — 
>

So in our case, we can begin by creating a new aggregate which is **Customer**.
The reason for an aggregate is that the business logic will be applied on the **Customer** aggregate, instead of each Entity holding the logic. An aggregate does not allow direct access to underlying entities. It is also common that multiple entities are needed to correctly represent data in real life, for instance, a Customer. It is a Person, but he/she can hold Products, and perform transactions.

An important rule in DDD aggregates is that they should only have one entity act as a **root entity**. What this means is that the reference of the root entity is also used to reference the aggregate. For our customer aggregate, this means that the **Person** ID is the unique identifier.

In the file, we will add a new struct named **Customer** and it will hold all needed entities to represent a Customer. Notice that all fields in the struct begins with *lower case letters*, this is a way in Go to make an object inaccessible from outside of the package the struct is defined in. This is done because an Aggregate should not allow direct access to the data. Neither does the struct define any tags for how the data is formatted such as json.