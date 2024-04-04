import React from 'react'
import {
    Collapse,
    Container,
    Nav,
    Navbar,
    NavbarToggler,
    NavItem,
} from 'reactstrap'
import { NavLink } from 'react-router-dom'
import Profile from './Profile'
import ThemeToggle from './ThemeToggle'
var csh = require('csh-material-bootstrap');

const NavBar: React.FunctionComponent = () => {
    const [isOpen, setIsOpen] = React.useState<boolean>(false)

    const toggle = () => {
        console.log(csh.ThemeToggle(null))
        setIsOpen(!isOpen)
    }

    return (
        <div>
            <Navbar color='primary' dark expand='lg' fixed='top'>
                <Container>
                    <NavLink to='/' className={'navbar-brand'}>
                        Rate A Professor
                    </NavLink>
                    <NavbarToggler onClick={toggle} />
                    <Collapse isOpen={isOpen} navbar>
                        <Nav navbar>
                            <NavItem>
                                <NavLink to='/' className='nav-link'>
                                    Home
                                </NavLink>
                            </NavItem>
                            {
                                // to add stuff to the navbar, add a NavItem tag with a NavLink to the route
                            }
                        </Nav>
                        <ThemeToggle/>
                        <Nav navbar className='ml-auto'>
                            <Profile />
                        </Nav>
                    </Collapse>
                </Container>
            </Navbar>
        </div>
    )
}

export default NavBar