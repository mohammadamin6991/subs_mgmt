this project, implements a simple microservice architecture consisting of the following services:

    Gateway Service
    Authentication Service
    Subscription Service
    svc_s3 (Example S3 Adapter Service)

Overview of the Workflow:

    User Authentication:
        Users interact with the Authentication Service to register or log in. Upon successful registration, the user is informed that the registration was successful, but no JWT token is issued at that time. The user must log in afterward to obtain a JWT token for subsequent requests.
        The Authentication Service also provides endpoints for renewing the JWT token (using a refresh token) and for revoking the token (during logout).

    Gateway Service:
        After obtaining the JWT token via login, users send requests to the Gateway Service, which acts as the entry point for routing to other backend services. The Gateway Service validates the JWT token by calling the Authentication Service to extract key user details (such as username and admin status).
        The Gateway Service enforces authorization rules, ensuring users with the appropriate privileges can access certain endpoints. It forwards the validated requests to the appropriate backend service (e.g., svc_s3) using a Golang reverse proxy.

    Service Adapters (e.g., svc_s3):
        svc_s3 is an example of an adapter service that manages interactions with MinIO, an S3-compatible object storage service. Each service adapter is responsible for interfacing with a specific backend service (e.g., storage, messaging, etc.), and in the future, additional service adapters can be added as needed.
        In the case of svc_s3, it interacts with MinIO to create a user, provision storage resources, and store the user’s service details.
        Initially, users do not have a service subscription. They can browse available subscription plans (CRUD operations can be performed by admin users). When a user selects a plan and requests a purchase, the adapter (like svc_s3) sends the selected plan's SKU, price, subscription duration, and username to the Subscription Service.

    Subscription Service:
        The Subscription Service processes the subscription request, generates an invoice, and provides a callback URL for payment processing.
        After the user completes the payment, the Subscription Service confirms the payment and sends a request to the appropriate service adapter (e.g., svc_s3) to create a new service record for the user.
        The service adapter (like svc_s3) then updates the user’s service details and triggers the necessary backend operations (e.g., creating storage resources in MinIO). This effectively activates the user's service.

    Viewing Service Details:
        Authenticated users can view their service details through the relevant service adapter (e.g., svc_s3), which retrieves relevant information from its database and communicates with the backend (e.g., MinIO) to fetch user-specific data as needed.

Key Technologies:

    JWT for user authentication and authorization.
    Refresh tokens for renewing JWTs and revoking tokens during logout.
    Golang reverse proxy for routing requests from the Gateway Service to backend services.
    MinIO (as an example) for object storage, with svc_s3 serving as an adapter. Additional service adapters can be added for other backend systems in the future.
