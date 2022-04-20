

describe("renders the contact page",()=>{
    it("renders correctly",()=>{
        cy.visit("http://localhost:3000/contact")
        cy.get('#contactContainer').should("exist")
    })
})