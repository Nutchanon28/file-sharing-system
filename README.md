# Docs

github.com/Nutchanon28/file-sharing-system

User Story: As a professor in a department in a certain university, I want a file sharing system with permission, so that I can share my files with people in my department only.

## Tech Stack and Why?

### PostGres (Jet)
- PostGres is selected by most developers
- Jet for type-safe query without ORM
- Not using ORM for...
  - better query efficiency
  - learning SQL in general (ORM is not agnostic; ORM is tied to specific programming languages or frameworks. SQL is universal)
- Alternatives: other DBMS

### GraphQL
- More efficient than REST but maybe a bit overkill
- Somewhat Popular
- Automate everything
- I want to learn it
- Alternatives: REST api

### gqlgen
- Popular go graphql library
- Alternatives: apollo-server

### Minio
- Free file storage but need to be deployed
- Popular, easy to use
- Alternatives: Amazon S3

## Common Problems

Problem: Jet command not found
Solution:
1. echo $GOPATH
2. ls $HOME/go/bin/jet
3. export PATH=$PATH:$HOME/go/bin

Problem: role *whatever* does not exist
Solution:
docker exec -it file-sharing-system-db-1 psql -h localhost -U rizzgod69 -d file-sharing-system

## References

### Why ORM might not be the best solution?
https://medium.com/the-existing/golang-why-orm-shouldnt-be-the-best-solution-199408536be

### Installing Go Jet
https://github.com/go-jet/jet?tab=readme-ov-file
