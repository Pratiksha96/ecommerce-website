describe("renders the contact page",()=>{
    it("renders correctly",()=>{
        cy.visit("http://localhost:3000/about")
        cy.get('.aboutSection').should("exist")
    })
})