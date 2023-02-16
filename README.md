# How to run syft? 

Install with `brew`: 
```bash
brew install syft
```

Run: 
```bash
syft . -c syft.yaml -o cyclonedx-json --file syft-bom.json
```

# How to run cdxgen? 

Install with `npm`: 
```bash
npm install -g @cyclonedx/cdxgen
```

Run: 
```bash
cdxgen -r  -o cdx-bom.json --deep
```
