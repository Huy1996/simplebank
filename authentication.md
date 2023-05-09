# PASETO

## JWT Singing Algorithms

- Symmetric digital signature algorithm
  - The same secret key is used to sign & verify token
  - For local use: interanl services, where the secret keu can be shared
  - HS256, HS384, HS512
    - HS256 = HMAC + SHA256
    - HMAC - Hash-based Message Authentication Code
    - SHA: Secure Hash Algorithm
    - 256/384/512: number of output bits

- Asymmetric digital signature algorithm
  - The private key is used to sign token
  - The public key is used to verify token
  - For public use: internal service signs token, but external service needs to verify it

### What's the problem of JWT

- Weak algorithms
  - Give developers too many algorithms to choose
  - Some algorithms are known to be vulnerable:
    - RSA PKCSv1.5: padding oracle attack
    - ECDSA: invalid-curve attack
- Trivial Forgery
  - set "alg" header to "none" to bypass security
  - set "alg" header to "HS256" while server normally verifies token with a RSA public key


## Platform-Agnostic Security Token [PASETO]

- Stronger algorithms
  - Developer don't have to choose the algorithm
  - Only need to select the version of PASETO
  - Each version has 1 strong cipher suite
  - Only 2 most recent PASTO versions are accepted
- Non-trivial Forgery
  - No more "alg" header or "none" algorithms
  - Everything is authenticated
  - Encrypted payload for local use \< symmetric key \>
  