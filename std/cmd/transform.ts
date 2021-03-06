import { Format, Overwrite } from '../index';
import * as host from '@jkcfg/std/internal/host'; // magic module
import * as param from '../param';
import { generate, File, GenerateParams } from './generate';
import { valuesFormatFromPath } from '../read';

type TransformFn = (value: any) => any | void;

const inputParams: GenerateParams = {
  stdout: param.Boolean('jk.transform.stdout', false),
  overwrite: param.Boolean('jk.transform.overwrite', false) ? Overwrite.Write : Overwrite.Err,
};

function transform(fn: TransformFn): void {

  function transformOne(obj: any): any {
    let txObj = fn(obj);
    txObj = (txObj === undefined) ? obj : txObj;
    return txObj;
  }

  const inputFiles = param.Object('jk.transform.input', {});
  const outputs = [];
  for (const path of Object.keys(inputFiles)) {
    const format = valuesFormatFromPath(path);
    outputs.push(host.read(path, { format }).then((obj): File => {
      switch (format) {
      case Format.YAMLStream:
      case Format.JSONStream:
        return {
          path,
          format,
          value: Array.prototype.map.call(obj, transformOne),
        };
      default:
        return { path, value: transformOne(obj) };
      }
    }));
  }
  generate(Promise.all(outputs), inputParams);
}

export default transform;
