# java-test-flow

### Test
```bash
faas template pull
faas template pull https://github.com/s8sg/faas-flow 
faas pull
faas build
faas deploy
echo "foo" | faas invoke date-flow
```
Change the `date-function` name in `stack.yml` to switch between function
```yml
date-function: "date-python"
```
