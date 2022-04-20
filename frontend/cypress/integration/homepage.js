describe("renders the home page",()=>{
    it("renders correctly",()=>{
        cy.visit("http://localhost:3000/")
        cy.get('#container').should("exist")
    })
})


it("gets headings",()=>{
    cy.visit("http://localhost:3000/")

    cy.contains("Featured Products")    
    })
    
//     describe('Get started', () => {
//         beforeEach(() => {
//           cy.visit("/signup") 
//         })
      
//         it('is redirected to the register page', () => {
//           cy.contains('Get Started')
//             .should('be.visible')
//             .click()
      
//           cy.url()
//             .should('be.equal', 'http://localhost:3000/signup')
//         })
//       })
    
      
//       describe('Register', () => {
//         beforeEach(() => {
//           cy.visit('http://localhost:3000/signup')
//         })
      
//         it('successfully registers', () => {
//           cy.get('.name')
//             .should('be.visible')
//             .type('Rohit')
//           cy.get('.email')
//             .should('be.visible')
//             .type('rohit.test@gmail.com')
//           cy.get('.password')
//             .should('be.visible')
//             .type('PWD')
//           cy.get('#nextbut').should('be.visible').click()
    
//             cy.get('.username').should('be.visible').type('testUser')
//             cy.get('.pass').should('be.visible').type('password')
    
//             cy.get('#nextbut').should('be.visible').click()
    
          
//             cy.url()
//             .should('be.equal', 'http://localhost:3000/signin')
    
    
//           })
//       })
    
    
     