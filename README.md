# mms-model
## Multi-Server Queueing Model (MMS) in Go

This code implements the Multi-Server Queueing Model (MMS) in the Go programming language. The MMS model is a mathematical tool used to analyze queueing systems with multiple servers, as in scenarios where multiple servers serve a queue of requests.

### How It Works

The `mmsModel` function calculates several important metrics for evaluating the performance of a queueing system with multiple servers. Here's how these metrics are calculated:

1. **Utilization Rate (rho):**
   The utilization rate is calculated as the ratio of the arrival rate of requests (lambda) to the average service rate per server (mu). This metric reflects how busy the servers are on average.

2. **Probability of an Empty System (p0):**
   The probability of having no customers in the system is calculated. This probability is used in the calculation of other metrics.

3. **Average Number of Customers in Queue (Lq):**
   The average number of customers in the queue at any given time is calculated. This provides insight into congestion in the queue and the efficiency of the system.

4. **Average Number of Customers in the System (L):**
   The average total number of customers in the system is calculated, including those in the queue and those being served by the servers.

5. **Average Time a Customer Spends in Queue (Wq):**
   The average time a customer spends in the queue before being served is calculated.

6. **Average Time a Customer Spends in the System (W):**
   The average time a customer spends in the entire system, including queue time and service time, is calculated.

### Using the Code

You can use the `mmsModel` function to calculate these metrics in your Go program. Simply provide the values of lambda (arrival rate), mu (service rate per server), and s (number of servers) as arguments, and the function will return the calculated metrics.

Make sure to include the `sumP0` and `factorial` functions in your program, as they are used by `mmsModel` to perform the necessary calculations.

```go
// Example usage
L, Lq, W, Wq := mmsModel(lambda, mu, s)
fmt.Printf("Average Number of Customers in the System (L): %.6f\n", L)
fmt.Printf("Average Number of Customers in Queue (Lq): %.6f\n", Lq)
fmt.Printf("Average Time a Customer Spends in the System (W): %.6f\n", W)
fmt.Printf("Average Time a Customer Spends in Queue (Wq): %.6f\n", Wq)
```

