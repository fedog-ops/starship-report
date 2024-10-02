import React from 'react'

const OfficerCard = ({officer}) => {
  return (<>
     <div>{officer.name}</div>
     <div>{officer.email}</div>
     <div>{officer.password}</div>
  
  </>
   
  )
}

export default OfficerCard