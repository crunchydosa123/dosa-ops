import React from 'react'

type Props = {
  children: React.ReactNode;
};

const AuthLayout = ({ children }: Props) => {
  return (
    <div className='w-screen h-screen flex'>
      

      <div className='w-full flex justify-center items-center'>{children}</div>
    </div>
  )
}

export default AuthLayout