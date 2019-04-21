import { requestAsPromise } from 'std_deferred';
import flatbuffers from 'flatbuffers';
import { __std } from '__std_generated';
import { Format } from 'std_write';

const Encoding = Object.freeze(__std.Encoding);

function uint8ToUint16Array(bytes) {
  return new Uint16Array(bytes.buffer, bytes.byteOffset, bytes.byteLength / 2);
}

const compose = (f, g) => x => f(g(x));
const stringify = bytes => String.fromCodePoint(...uint8ToUint16Array(bytes));

// read requests the path and returns a promise that will be resolved
// with the contents at the path, or rejected.
function read(path, opts = {}) {
  const { encoding = Encoding.JSON, format = Format.FromExtension, module } = opts;

  const builder = new flatbuffers.Builder(512);
  const pathOffset = builder.createString(path);
  let moduleOffset = 0;
  if (module !== undefined) {
    moduleOffset = builder.createString(module);
  }
  __std.ReadArgs.startReadArgs(builder);
  __std.ReadArgs.addPath(builder, pathOffset);
  __std.ReadArgs.addEncoding(builder, encoding);
  __std.ReadArgs.addFormat(builder, format);
  if (module !== undefined) {
    __std.ReadArgs.addModule(builder, moduleOffset);
  }
  const argsOffset = __std.ReadArgs.endReadArgs(builder);
  __std.Message.startMessage(builder);
  __std.Message.addArgsType(builder, __std.Args.ReadArgs);
  __std.Message.addArgs(builder, argsOffset);
  const messageOffset = __std.Message.endMessage(builder);
  builder.finish(messageOffset);

  let tx = bytes => bytes;
  switch (encoding) {
  case Encoding.String:
    tx = stringify;
    break;
  case Encoding.JSON:
    tx = compose(JSON.parse, stringify);
    break;
  default:
    break;
  }

  return requestAsPromise(() => V8Worker2.send(builder.asArrayBuffer()), tx);
}

export {
  Encoding,
  read,
};
