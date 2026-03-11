import { Card } from '../../components/ui/card'

const DbMain = () => {
  return (
    <div className='w-full border bg-blue-300'>
      <div className='p-5 flex flex-col'>
        
        <div className='text-2xl font-semibold'>Dashboard</div>
        <div className='w-full my-6 h-36 flex gap-3'>
          <div className='h-full w-1/3 '><Card className='h-full'/></div>
          <div className='h-full w-1/3 '><Card className='h-full'/></div>
          <div className='h-full w-1/3 '><Card className='h-full'/></div>
        </div>

        <div className='h-110 flex gap-4'>
          <Card className='w-2/3 h-full' />
          <Card className='w-1/3 h-full' />
        </div>
      </div>
    </div>
  )
}

export default DbMain