describe('Main', () => {
    beforeEach(() => {
      cy.visit("/") 
    })
  
    it('renders main page correctly', () => {
      
        
      cy.url()
        .should('be.equal', 'http://localhost:3000/')
    })


    it('has 8 cards', () => {
       

       cy.get('.card').its('length').should('be.eq', 8)
      })


      it('renders product page on click', () => {
       

        cy.contains('Products').click()

        cy.url()
        .should('be.equal', 'http://localhost:3000/product')

        cy.visit('/')


       })

       it('renders single product page on click on product', () => {
       

        cy.contains('product').click()

        cy.url()
        .should('be.equal', 'http://localhost:3000/product/:id')

        cy.visit('/')


       })

  })

