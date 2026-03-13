import AuthLayout from './AuthLayout'
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from '../../components/ui/card'
import { Label } from '../../components/ui/label'
import { Input } from '../../components/ui/input'
import { Button } from '../../components/ui/button'

const Signup = () => {
  return (
    <AuthLayout>
      <div className="flex items-center justify-center h-full">

        <Card className="w-[380px]">

          <CardHeader>
            <CardTitle className="text-xl">Signup</CardTitle>
          </CardHeader>

          <CardContent className="flex flex-col gap-4">

            <div className="flex flex-col gap-2">
              <Label>Name</Label>
              <Input type="email" placeholder="you@example.com" />
            </div>

            <div className="flex flex-col gap-2">
              <Label>Email</Label>
              <Input type="email" placeholder="you@example.com" />
            </div>

            <div className="flex flex-col gap-2">
              <Label>Password</Label>
              <Input type="password" placeholder="••••••••" />
            </div>

            <div className="flex flex-col gap-2">
              <Label>Confirm Password</Label>
              <Input type="password" placeholder="••••••••" />
            </div>

          </CardContent>

          <CardFooter>
            <Button className="w-full">Sign In</Button>
          </CardFooter>

        </Card>

      </div>
    </AuthLayout>
  )
}

export default Signup