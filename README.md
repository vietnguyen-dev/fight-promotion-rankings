A ranking systsem for fight promotions

Smaller fight promotions in the United States and Canada do not have rankings on their websites. Promotion rankings is a website that allows fight promotions to compare their ranks to be able to cross promote

Admin App + Public Websites - React TypeScript Vite with SSR
Backend - Go
Database - PG
Hosting - AWS EC2 + Docker Compose

Services
User Management - Clerk
image storage - aws s3


admin app features
1. users should be able to log in and out, one user in charge of one fight promotion
2. while logged in users should be able to perform CRUD operations on a list of fights, fighters, weight classes for their promotion
    - 4 pages, events, athletes, rankings
    - weight classes
        - CRUD FOR weight classes
    - events
        - list of events that are dropdowns to see all fights within
        - all fights are dropdowns where user can update weight class, winners, losers, result, time 
    - athletes
        - CRUD for athletes
    - rankings
        - users will be able to choose from athletes from their roster and be able to order their rankings to be shown on the public website

public websites feature
1. promotion rankings should be publicly be available on this-domain.com/promotion-code
    - rankings by weight class and p4p rankings





