import React from "react";
 
import "./style.css";

export const MyPlugin = () => {
  return (
    <div id="webcrumbs"> 
    	<div className="w-full h-screen p-8 bg-neutral-50 rounded-lg shadow-lg flex flex-col gap-6 relative">
    	  <img 
    	    src="https://cdn.webcrumbs.org/assets/images/ask-ai/bgs/18.svg" 
    	    className="absolute top-0 left-0 w-full h-full object-cover opacity-5" 
    	    alt="" 
    	  />
    	
    	  <h1 className="text-3xl font-title text-primary relative z-10">Enterprise Speech to Text Converter</h1>
    	
    	  <textarea
    	    className="w-full rounded-md border border-neutral-300 p-4 text-lg bg-neutral-50 focus:outline-none focus:ring-2 focus:ring-primary relative z-10"
    	    rows="10"
    	    placeholder="Your transcribed speech will be displayed here..."
    	  ></textarea>
    	
    	  <div className="flex justify-center gap-10 mt-6 relative z-10">
    	    <button className="rounded-md bg-primary text-white h-[50px] w-[150px] text-lg font-semibold transition-transform duration-200 ease-in-out hover:scale-105">
    	      Start
    	    </button>
    	    <button className="rounded-md bg-primary text-white h-[50px] w-[150px] text-lg font-semibold transition-transform duration-200 ease-in-out hover:scale-105">
    	      Stop
    	    </button>
    	    <button className="rounded-md bg-primary text-white h-[50px] w-[150px] text-lg font-semibold transition-transform duration-200 ease-in-out hover:scale-105">
    	      Reset
    	    </button>
    	  </div>
    	
    	  <div className="mt-8 flex flex-col gap-4">
    	    <details className="border rounded-md p-4 bg-neutral-100 relative z-10 shadow-sm">
    	      <summary className="font-semibold cursor-pointer">Advanced Settings</summary>
    	      <div className="mt-4">
    	        <label className="block font-medium mb-2">Language:</label>
    	        <select className="w-full rounded-md border border-neutral-300 p-2 bg-neutral-50 focus:outline-none focus:ring-2 focus:ring-primary text-sm">
    	          <option>English (US)</option>
    	          <option>English (UK)</option>
    	          <option>Spanish</option>
    	          <option>French</option>
    	        </select>
    	
    	        <label className="block font-medium mt-4 mb-2">Microphone Sensitivity:</label>
    	        <input type="range" min="0" max="100" className="w-full cursor-pointer" />
    	      </div>
    	    </details>
    	
    	    <details className="border rounded-md p-4 bg-neutral-100 relative z-10 shadow-sm">
    	      <summary className="font-semibold cursor-pointer">Output Format</summary>
    	      <div className="mt-4">
    	        <label className="block font-medium mb-2">Choose Format:</label>
    	        <select className="w-full rounded-md border border-neutral-300 p-2 bg-neutral-50 focus:outline-none focus:ring-2 focus:ring-primary text-sm">
    	          <option>Text (.txt)</option>
    	          <option>Word (.docx)</option>
    	          <option>PDF (.pdf)</option>
    	        </select>
    	      </div>
    	    </details>
    	  </div>
    	
    	  <img 
    	    src="https://cdn.webcrumbs.org/assets/images/ask-ai/gradients/g4.png" 
    	    className="absolute bottom-0 right-0 w-[120px] h-[120px] object-cover rounded-full"
    	    alt="" 
    	  />
    	</div> 
    </div>
  )
}

