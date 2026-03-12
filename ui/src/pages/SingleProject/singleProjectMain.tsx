import { useState } from "react"

import { Card, CardHeader, CardTitle, CardContent } from "../../components/ui/card"
import { Button } from "../../components/ui/button"

import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter
} from "../../components/ui/dialog"

import { Input } from "../../components/ui/input"
import { Label } from "../../components/ui/label"
import { Badge } from "../../components/ui/badge"

type Service = {
  name: string
  url: string
  apiKey: string
  health: "healthy" | "down"
}

const SingleProjectMain = () => {
  const [services, setServices] = useState<Service[]>([])

  const [name, setName] = useState("")
  const [url, setUrl] = useState("")
  const [apiKey, setApiKey] = useState("")

  const addService = () => {
    if (!name || !url) return

    setServices([
      ...services,
      {
        name,
        url,
        apiKey,
        health: "healthy"
      }
    ])

    setName("")
    setUrl("")
    setApiKey("")
  }

  return (
    <div className="p-6 flex flex-col gap-6">

      {/* Header */}
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-semibold">Project Services</h1>

        <Dialog>
          <DialogTrigger asChild>
            <Button>Add Service</Button>
          </DialogTrigger>

          <DialogContent>
            <DialogHeader>
              <DialogTitle>Add New Service</DialogTitle>
            </DialogHeader>

            <div className="flex flex-col gap-4 py-2">

              <div>
                <Label>Service Name</Label>
                <Input value={name} onChange={(e)=>setName(e.target.value)} />
              </div>

              <div>
                <Label>Service URL</Label>
                <Input value={url} onChange={(e)=>setUrl(e.target.value)} />
              </div>

              <div>
                <Label>API Key</Label>
                <Input value={apiKey} onChange={(e)=>setApiKey(e.target.value)} />
              </div>

            </div>

            <DialogFooter>
              <Button onClick={addService}>Create Service</Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      </div>

      {/* Services Grid */}
      <div className="grid grid-cols-3 gap-4">

        {services.map((service, i) => (
          <Card key={i} className="hover:shadow-md">

            <CardHeader className="flex flex-row justify-between items-center">
              <CardTitle className="text-lg">{service.name}</CardTitle>

              <Badge
                variant={service.health === "healthy" ? "default" : "destructive"}
              >
                {service.health}
              </Badge>

            </CardHeader>

            <CardContent className="text-sm text-muted-foreground">
              {service.url}
            </CardContent>

          </Card>
        ))}

        {services.length === 0 && (
          <div className="text-muted-foreground">
            No services added yet.
          </div>
        )}

      </div>

    </div>
  )
}

export default SingleProjectMain