describe('signs ',()=>{
    beforeEach(() => {
      cy.visit('http://localhost:3000/login')
    })

    it('signs in correctly',()=>{

      cy.get('.loginEmail').should('be.visible').type('a@ufl.edu')
    cy.get('.loginPassword').should('be.visible').type('123123123')

      cy.get('.loginBtn').click()

    })
    

  })