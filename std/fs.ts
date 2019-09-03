/**
 * @module std/fs
 */

import { RPCSync } from './internal/rpc';
import { valueFromUTF8Bytes } from './internal/data';

export class FileInfo {
  name: string;
  path: string;
  isdir: boolean;

  constructor(n: string, p: string, d: boolean) {
    this.name = n;
    this.path = p;
    this.isdir = d;
  }
}

export class Directory {
  name: string;
  path: string;
  files: FileInfo[];

  constructor(n: string, p: string, files: FileInfo[]) {
    this.name = n;
    this.path = p;
    this.files = files;
  }
}

export interface InfoOptions {
  module?: string;
}

export function info(path: string, options: InfoOptions = {}): FileInfo {
  const { module = '' } = options;
  const response = RPCSync('std.fileinfo', path, module);
  const { name: n, path: p, isdir: d } = valueFromUTF8Bytes(response);
  return new FileInfo(n, p, d);
}

export interface DirOptions {
  module?: string;
}

export function dir(path: string, options: DirOptions = {}): Directory {
  const { module = '' } = options;
  const response = RPCSync('std.dir', path, module);
  const { name: n, path: p, files: fs } = valueFromUTF8Bytes(response);
  const infos = [];
  for (const f of fs) {
    const { name: infoname, path: infopath, isdir: infoisdir } = f;
    infos.push(new FileInfo(infoname, infopath, infoisdir));
  }
  return new Directory(n, p, infos);
}
