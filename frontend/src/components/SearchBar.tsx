import React, { useState, ChangeEvent, FormEvent } from 'react';
import './SearchBar.css'

interface SearchBarProps {
  onSearch: (query: string) => void;
}

const SearchBar: React.FC<SearchBarProps> = ({ onSearch }) => {
  const [query, setQuery] = useState<string>('');

  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    setQuery(event.target.value);
  };

  const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    onSearch(query);
  };

  // return(
  //   <form onSubmit={handleSubmit} className="form-inline my-1 my-lg-0">
  //     <input className="form-control mr-sm-1" type="search" placeholder="Search" aria-label="Search" value={query} onChange={handleChange}/>
  //     <button className="btn btn-outline-success my-1 my-sm-0" type="submit">Search</button>
  //   </form>
  // )

  return (
    
    <form onSubmit={handleSubmit} className='search-bar'>
    <input 
      type="search" 
      name="search" 
      pattern=".*\S.*" 
      required 
      value={query}
      onChange={handleChange}
      />
    <button className="search-btn" type="submit"/>
  </form>
  );
};

export default SearchBar;
