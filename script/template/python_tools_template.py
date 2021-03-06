#!/usr/bin/env python2.7
#_*_coding:utf8_*_

from __future__ import with_statement
from optparse import OptionParser
from os import system as sh
from time import strftime
import base64, random
import sys, os, time, commands, re, string
import getpass,subprocess
import commands;


cmds = ["help"]

def addSudoToUser():
    #try check sudo pri
    sh('sudo usermod -aG sudo ' + getpass.getuser());

def usage(option, opt, value, parser):
    parser.print_help()
    print "Arguments:"
    for item in cmds:
        cmd_help = eval("%s.__doc__" % item)
        print cmd_help.split("\n")[0]

def get_lazy_cmd(cmd_name):
    if cmd_name in cmds: return cmd_name
    lazy_cmds = [cmd for cmd in cmds if cmd.startswith(cmd_name)]
    if lazy_cmds == []:
        print "匹配不到命令:%s" % cmd_name
        return None
    if len(lazy_cmds) > 1:
        print "命令不能唯一匹配：", lazy_cmds
        return None
    return lazy_cmds[0]

def help(argv):
    """  help\t\thelp cmd显示命令的详细帮助

  ./syncImagess help db (显示restart的详细帮助信息)
"""
    if argv == []: 
        cmd_list = cmds[:]
    elif argv[0] in cmds:
        cmd_list = [cmds[cmds.index(argv[0])]]
    else:
        cmd_list = ["help"]
    for cmd in cmd_list:
        cmd_help = eval("%s.__doc__" % cmd)
        print cmd_help

def run():
    try:     
        #init
        addSudoToUser();
        
        parser = OptionParser(usage=u"./tools %s [options]" % "|".join(cmds), add_help_option=False)
        parser.set_defaults(run_mode=None,
                            guard_mode=False,
                            run_for_debug=False,
                            ignore_error=False)
        parser.add_option("-h", "--help", help=u"显示这些帮助信息", action="callback", callback=usage)
        parser.add_option("-c", help=u"指定配置文件的位置", dest="config_file")

        options, args = parser.parse_args()
        
        if len(args) >= 1:
            cmd = get_lazy_cmd(args[0])
            if not cmd: return 1
            if cmd != args[0]:
                print "%s匹配成%s" % (args[0], cmd)
            exec("%s(%s)" % (cmd, args[1:]))
        else:
            print "输入 %s -h 查看帮助." % sys.argv[0]
    except KeyboardInterrupt:
        print "用户中断操作！"
    except SystemExit, error_no:
        return error_no
    except Exception, e:
        print e
        
    return 1

if __name__ == "__main__":
    sys.exit(run())
