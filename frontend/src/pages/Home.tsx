import { useState }  from 'react'
import SearchBar from '../components/SearchBar'
import './Home.css'


const Home = () => {
    // important hooks
    // const { accessTokenPayload } = useOidcAccessToken()   // this contains the user info in raw json format
    // const userInfo = accessTokenPayload as UserInfo       // 
    // const { idToken, idTokenPayload } = useOidcIdToken()  // this is how you get the users id token
    // const { login, logout, isAuthenticated } = useOidc()  // this gets the functions to login and logout and the logout state

    const [searchResults, setSearchResults] = useState<string[]>([]);

    const handleSearch = (query: string) => {
        // Perform search operation here (e.g., fetch data from API)
        console.log('Searching for:', query);
        // Placeholder for search results
        const searchResults: string[] = []; // Replace this with your actual search results
        // Update searchResults state with search results
        setSearchResults(searchResults);
    };

    return (
        <div className='home'>
            <h2 className="display-3">Find a professor!</h2>
            <SearchBar onSearch={handleSearch} />
      {/* Display search results */}
      <ul>
        {searchResults.map((result, index) => (
          <li key={index}>{result}</li>
        ))}
      </ul>
        </div>
    )
}

export default Home